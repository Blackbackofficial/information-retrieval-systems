#include <fstream>
#include <vector>
#include "Record.h"
#include "Logger.h"

/// Simulate tape in polyphase merge polyphase_sort. Stores input and output handlers to files, provide data block read/write mode
class Tape {
private:
    const char VALUE_SEP = ',';
    const char RECORD_SEP = ';';
    Logger *log; // Logger passed in constructor
    const char *filepath; // Tape path
    int block_size; // Size of read block
    int block_index; // Point at char in block/buffer
    std::vector<char> block; // Buffer for read/save data
    std::ifstream file_in_stream; // Input handler to given file
    std::ofstream file_out_stream; // Output handler to given file

    /// Initiate tape to read from file
    void init_read() {
        close_tape();
        file_in_stream.open(filepath, std::ifstream::binary);
        get_block();
    }

    /// Initiate tape to save to file
    void init_save() {
        close_tape();
        file_out_stream.open(filepath, std::ofstream::binary);
    }

    /// Increment block_index and check if next block/page needs to be read/saved
    void inc_block() {
        if (++block_index == block_size) {
            block_index = 0;
            if (file_in_stream.is_open() and file_in_stream.eof())
                block_size = 0;
            else if (file_in_stream.is_open())
                get_block();
            else if (file_out_stream.is_open())
                save_block();
        }
    }

    /// Read next block of data up to block size
    void get_block() {
        file_in_stream.read(block.data(), block.size());
        block_size = file_in_stream.gcount();
        log->inc_inout_number();
    }

    /// Save next block of data up to block size
    void save_block() {
        file_out_stream.write(block.data(), block_size);
        log->inc_inout_number();
    }

    /// Close all connection in tape
    void close_tape() {
        if (file_out_stream.is_open()) {
            // Save out stream
            block_size = block_index;
            save_block();
            file_out_stream.close();
        } else if (file_in_stream.is_open()) {
            file_in_stream.close();
        }
        // Reset block
        block_index = 0;
        block_size = block.size();
    }

public:
    /**
     * Constructor of tape:
     *  set filename,
     *  initiate data block
     * @param filepath to tape
     * @param block_size in bytes
     */
    Tape(const char *filepath, int block_size, Logger &log) {
        this->log = &log;
        this->filepath = filepath;
        this->block_size = block_size;
        this->block_index = 0;
        this->block.resize(this->block_size, 0);

        this->file_in_stream.close();
        this->file_out_stream.close();
    }

    /**
     * Initiate read mode if necessary and read values from file
     * @return next record object as smart ptr or nullptr if file ended
     */
    auto get_record() -> std::shared_ptr<Record> {
        // Create input handler to file
        if (!file_in_stream.is_open()) init_read();

        // Return dummy record
        if (file_in_stream.eof() and block_index == block_size) return nullptr;

        std::string tmp_height, tmp_base;
        // Read first value to tmp string from buffer/block
        while (block[block_index] != VALUE_SEP) {
            tmp_height.push_back(block[block_index]);
            inc_block();
        }
        inc_block();

        // Read second value to tmp vector from buffer/block
        while (block[block_index] != RECORD_SEP) {
            tmp_base.push_back(block[block_index]);
            inc_block();
        }
        inc_block();

        // Cast string to int
        auto height = std::stoi(tmp_height);
        auto base = std::stoi(tmp_base);

        return std::make_shared<Record>(height, base);
    }

    /**
     * Save record to file
     * @param r record to save in file
     */
    void save_record(const std::shared_ptr<Record> &r) {
        // Create output handler to file
        if (!file_out_stream.is_open()) init_save();

        auto tmp_height = std::to_string(r->get_height());
        auto tmp_base = std::to_string(r->get_base());

        // Save first value from tmp str to buffer/block
        for (auto i : tmp_height) {
            block[block_index] = i;
            inc_block();
        }
        block[block_index] = VALUE_SEP;
        inc_block();

        // Save second value from tmp str to buffer/block
        for (auto i : tmp_base) {
            block[block_index] = i;
            inc_block();
        }
        block[block_index] = RECORD_SEP;
        inc_block();
    }

    /// Print content of tape after distribute
    void print_tape() {
        if (file_out_stream.is_open()) {
            close_tape();
            log->print_file_debug(filepath);
        }
    }

    ~Tape() { close_tape(); }
};
