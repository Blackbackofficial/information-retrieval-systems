#include <iostream>
#include <memory>

/// Storage for single record from file that represents values of triangle (height and base)
class Record {
private:
    int height; // height of triangle
    int base; // base of triangle

public:
    /**
     * Constructor of record
     * @param height of triangle
     * @param base of triangle
     */
    Record(int height, int base) {
        this->height = height;
        this->base = base;
    }

    auto get_height() const -> int {
        return height;
    }

    auto get_base() const -> int {
        return base;
    }

    /**
     * @param r other record to compare
     * @return true if current record (triangle) has bigger or equal area
     */
    auto is_area_bigger(std::shared_ptr<Record> &r) const -> bool {
        return (long long) r->get_base() * (long long) r->get_height() <= (long long) get_base() * (long long) get_height();
    }

    ~Record() = default;
};
