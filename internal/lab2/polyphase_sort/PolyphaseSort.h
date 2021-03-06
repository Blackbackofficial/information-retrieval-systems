#include "Tape.h"
#include <filesystem>
#include <tuple>

/// Sort given file by polyphase merge method and save it to given output
class PolyphaseSort {
private:
    Logger* log; // Logger passed in constructor
    const char *out_filepath; // Path to sorted output file
    int tape_distribution_id; // Number of tape to Distribute
    int tape_offset; // Offset to read current tape
    int dummy_series; // Number of empty series from init phase
    std::unique_ptr<Tape> input_tape; // Input tape from file
    std::unique_ptr<Tape> tapes[3]; // Three tapes to distribute sorted records
    std::shared_ptr<Record> records[3]; // Current read records from files

    /// Init distribution from input into two tapes
    void init_phase() {
        // Init values
        int fib0 = 0, fib1 = 0, series_counter = 0;
        auto tape_save_id = get_tape_index();

        // Init record
        std::shared_ptr<Record> r, prev_r = input_tape->get_record();
        tapes[tape_save_id]->save_record(prev_r);

        // Read all records from tape
        while ((r = input_tape->get_record())) {
            // Series ended
            if (!r->is_area_bigger(prev_r)) ++series_counter;

            if (series_counter >= fib1) {
                // Save previous record in tape
                records[tape_save_id] = prev_r;

                // Change tape
                change_tape_offset();
                tape_save_id = get_tape_index();

                // Does series merge with previous record in tape
                series_counter = (records[tape_save_id] and r->is_area_bigger(records[tape_save_id])) ? -1 : 0;

                // Calc Fibonacci
                std::tie(fib0, fib1) = (fib1 == 0) ? std::make_tuple(0, 1) : std::make_tuple(fib1, fib0 + fib1);
            }

            // Save record to tape
            tapes[tape_save_id]->save_record(r);
            prev_r = r;
        }
        ++series_counter;
        dummy_series = fib1 - series_counter;
        change_tape_offset();

        // Clear
        log->print_dummy_series_info(dummy_series);
        log->print_series_info(fib1, fib0 + series_counter);
        input_tape = nullptr;
        records[0] = nullptr;
        records[1] = nullptr;
        records[2] = nullptr;
    }

    /// Remove dummy series from tape
    void dummy_phase() {
        // Init values
        int input_ids[2] = {(tape_distribution_id + 1) % 3, (tape_distribution_id + 2) % 3};

        // Init records
        if (!records[input_ids[0]]) records[input_ids[0]] = tapes[input_ids[0]]->get_record();
        if (!records[input_ids[1]]) records[input_ids[1]] = tapes[input_ids[1]]->get_record();
        std::shared_ptr<Record> r;

        // Read dummy series
        while (dummy_series) {
            // Save record from normal series
            tapes[tape_distribution_id]->save_record(records[input_ids[tape_offset]]);

            // Read next record from same tape
            r = tapes[input_ids[tape_offset]]->get_record();

            // Does series end
            if (!r or !r->is_area_bigger(records[input_ids[tape_offset]]))
                --dummy_series;
            records[input_ids[tape_offset]] = r;
        }

        // Dummy distribution end tape
        if (!records[input_ids[0]] or !records[input_ids[1]]) {
            // Switch distribution tape
            log->inc_phase_number();
            tapes[tape_distribution_id]->print_tape();
            if (records[0] or records[1] or records[2]) tape_distribution_id = input_ids[tape_offset];
            records[tape_distribution_id] = nullptr;
        } else phase();
    }

    /// Distribute and merge values from two tapes
    void phase() {
        // Init values
        log->inc_phase_number();
        log->print_distribution_tape_number_debug(tape_distribution_id + 1);
        bool is_series_end[2] = {false, false}, is_tape_end = false;
        int input_ids[2] = {(tape_distribution_id + 1) % 3, (tape_distribution_id + 2) % 3};

        // Init records
        if (!records[input_ids[0]]) records[input_ids[0]] = tapes[input_ids[0]]->get_record();
        if (!records[input_ids[1]]) records[input_ids[1]] = tapes[input_ids[1]]->get_record();
        tape_offset = (int) (records[input_ids[0]]->is_area_bigger(records[input_ids[1]]));
        std::shared_ptr<Record> r;

        // Read one tape and all series up to end
        while (!is_tape_end or !(is_series_end[0] and is_series_end[1])) {
            // Init new series
            if (!is_tape_end and is_series_end[0] and is_series_end[1]) {
                std::tie(is_series_end[0], is_series_end[1]) = std::make_tuple(false, false);
                tape_offset = (int) (records[input_ids[0]]->is_area_bigger(records[input_ids[1]]));
            }

            tapes[tape_distribution_id]->save_record(records[input_ids[tape_offset]]);

            // End of tape
            if (!(r = tapes[input_ids[tape_offset]]->get_record())) {
                is_tape_end = true;
                is_series_end[tape_offset] = true;
                records[input_ids[tape_offset]] = nullptr;
                if (!(is_series_end[0] and is_series_end[1]))
                    change_tape_offset();
                continue;
            }
            // Handle record
            if (!r->is_area_bigger(records[input_ids[tape_offset]])) {
                // End of series in current tape
                records[input_ids[tape_offset]] = r;
                is_series_end[tape_offset] = true;
                change_tape_offset();
            } else if (auto t2 = (tape_offset + 1) % 2; !is_series_end[t2] and r->is_area_bigger(records[input_ids[t2]])) {
                // Swap second tape if available
                records[input_ids[tape_offset]] = r;
                change_tape_offset();
            } else {
                // Stay on current tape
                records[input_ids[tape_offset]] = r;
            }
        }
        // Switch distribution tape
        tapes[tape_distribution_id]->print_tape();
        if (records[0] or records[1] or records[2]) tape_distribution_id = input_ids[tape_offset];
        records[tape_distribution_id] = nullptr;
    }

    /// tape_offset = (tape_offset + 1) % 2;
    void change_tape_offset() {
        tape_offset = (tape_offset + 1) % 2;
    }

    /// (tape_distribution_id + tape_offset + 1) % 3;
    auto get_tape_index() -> int {
        return (tape_distribution_id + tape_offset + 1) % 3;
    }

public:
    /**
     * Constructor of polyphase_sort class
     * @param in_filepath path to file to polyphase_sort
     * @param out_filepath path to save sorted file
     * @param block_size
     */
    PolyphaseSort(const char *in_filepath, const char *out_filepath, int block_size, Logger &log) {
        this->log = &log;
        this->out_filepath = out_filepath;
        this->input_tape = std::make_unique<Tape>(in_filepath, block_size, log);

        //init tapes
        this->tapes[0] = std::make_unique<Tape>("/Users/iva.chernov/GolandProjects/information-retrieval-systems/internal/lab2/polyphase_sort/tapes/tape1", block_size, log);
        this->tapes[1] = std::make_unique<Tape>("/Users/iva.chernov/GolandProjects/information-retrieval-systems/internal/lab2/polyphase_sort/tapes/tape2", block_size, log);
        this->tapes[2] = std::make_unique<Tape>("/Users/iva.chernov/GolandProjects/information-retrieval-systems/internal/lab2/polyphase_sort/tapes/tape3", block_size, log);
        this->tape_distribution_id = 2;
        this->tape_offset = 0;
        this->dummy_series = 0;

        //clear records
        this->records[0] = nullptr;
        this->records[1] = nullptr;
        this->records[2] = nullptr;
    }

    void sort() {
        // Run all phases
        init_phase();
        dummy_phase();
        while (records[0] or records[1] or records[2]) phase();

        // Close all tapes
        tapes[0] = nullptr;
        tapes[1] = nullptr;
        tapes[2] = nullptr;

        // Save output file
        auto output_tape = "/Users/iva.chernov/GolandProjects/information-retrieval-systems/internal/lab2/polyphase_sort/tapes/tape" + std::to_string(tape_distribution_id + 1);
        if (std::rename(output_tape.c_str(), out_filepath)) {
            log->print_save_output_error(out_filepath);
            exit(errno);
        }
    }

    ~PolyphaseSort() = default;
};