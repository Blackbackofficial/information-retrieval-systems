#include <iostream>
#include <fstream>
#include <cmath>

class Logger {
private:
    const char* logger_name;
    int logger_lever;
    int phase_number = 0;
    int inout_number = 0;

public:
    /**
     * @param logger_name
     * @param logger_lever
     */
    Logger(const char *logger_name, int logger_lever) {
        this->logger_name = logger_name;
        this->logger_lever = logger_lever;
    }

    ~Logger() {}

    void inc_inout_number() {
        inout_number++;
    }

    void inc_phase_number() {
        phase_number++;
    }

    void print_line() {
        if (logger_lever >= 1) {
            std::cout << "--------------------------------------------------------" << std::endl;
        }
    }

    void print_save_output_error(const char *filepath) {
        std::cout << "[ERROR] Error while saving tape to given output path: " << filepath << std::endl;
    }

    /**
     * @param dummy_series_number
     */
    void print_dummy_series_info(int dummy_series_number) {
        if (logger_lever >= 1) {
            std::cout << "[INFO] Number of dummy series: " << dummy_series_number << std::endl;
        }
    }

    /**
     * @param fib0
     * @param fib1
     */
    void print_series_info(int fib0, int fib1) {
        if (logger_lever >= 1) {
            std::cout << "[INFO] Number of series on tapes: " << fib0 << " | " << fib1 << std::endl;
        }
    }

    /**
     * Write in output given number of phases
     * @param phase_number
     */
    void print_phase_number() {
        if (logger_lever >= 1) {
            std::cout << "[INFO] Number of phases in polyphase_sort: " << phase_number << std::endl;
        }
    }

    /**
     * Write in output given number of read/write operations
     * @param inout_number
     */
    void print_inout_number() {
        if (logger_lever >= 1) {
            std::cout << "[INFO] Number of read/write in polyphase_sort: " << inout_number << std::endl;
        }
    }

    /**
     * Write in output calculated, expected number of series
     * @param series_number number of series saved in file
     */
    void print_theoretic_phase_number(long long records_number) {
        if (logger_lever >= 1) {
            std::cout << "[INFO] Theoretical number of phases: " << std::ceil(1.45 * log2(records_number / 2)) << std::endl;
        }
    }

    /**
     * Write in output calculated, expected number of write/read operation to files
     * @param record_number number of records in file
     * @param block_size size of reading/write block
     */
    void print_theoretic_inout_number(long long records_number, long long file_size, int block_size) {
        if (logger_lever >= 1) {
            std::cout << "[INFO] Theoretical number of read/write: ";
            std::cout << (int) std::ceil(2 * file_size * (1.04 * log2(records_number / 2) + 1) / block_size) << std::endl;
        }
    }

    /**
     * Print content of file
     * @param filepath
     */
    void print_file_debug(const std::string filepath) {
        std::ifstream f(filepath);

        if (logger_lever >= 2 and f.is_open()) {
            std::cout << "[DEBUG] Content of " << filepath << ":" << std::endl;
            std::cout << f.rdbuf() << std::endl;
            f.close();
        }
    }

    /**
     * @param number of tape
     */
    void print_distribution_tape_number_debug(int d) {
        if (logger_lever >= 2) {
            std::cout << "[DEBUG] Distribute series to tape" << d << std::endl;
        }
    }
};
