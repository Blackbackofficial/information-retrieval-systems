#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

static int NUM_OF_TAPES = 4;

// MergeSort
void merge(int array[], int const left, int const mid, int const right) {
    auto const subArrayOne = mid - left + 1;
    auto const subArrayTwo = right - mid;

    auto *leftArray = new int[subArrayOne], *rightArray = new int[subArrayTwo];

    // Copy data to temp arrays leftArray[] and rightArray[]
    for (auto i = 0; i < subArrayOne; i++) {
        leftArray[i] = array[left + i];
    }
    for (auto j = 0; j < subArrayTwo; j++) {
        rightArray[j] = array[mid + 1 + j];
    }

    auto indexOfSubArrayOne = 0, // Initial index of first sub-array
    indexOfSubArrayTwo = 0; // Initial index of second sub-array
    int indexOfMergedArray = left; // Initial index of merged array

    // Merge the temp arrays back into array[left..right]
    while (indexOfSubArrayOne < subArrayOne && indexOfSubArrayTwo < subArrayTwo) {
        if (leftArray[indexOfSubArrayOne] <= rightArray[indexOfSubArrayTwo]) {
            array[indexOfMergedArray] = leftArray[indexOfSubArrayOne];
            indexOfSubArrayOne++;
        } else {
            array[indexOfMergedArray] = rightArray[indexOfSubArrayTwo];
            indexOfSubArrayTwo++;
        }
        indexOfMergedArray++;
    }
    // Copy the remaining elements of
    // left[], if there are any
    while (indexOfSubArrayOne < subArrayOne) {
        array[indexOfMergedArray] = leftArray[indexOfSubArrayOne];
        indexOfSubArrayOne++;
        indexOfMergedArray++;
    }
    // Copy the remaining elements of
    // right[], if there are any
    while (indexOfSubArrayTwo < subArrayTwo) {
        array[indexOfMergedArray] = rightArray[indexOfSubArrayTwo];
        indexOfSubArrayTwo++;
        indexOfMergedArray++;
    }
}

void mergeSort(int array[], int const begin, int const end) {
    if (begin >= end) {
        return;
    }

    auto mid = begin + (end - begin)/2;
    mergeSort(array, begin, mid);
    mergeSort(array, mid + 1, end);
    merge(array, begin, mid, end);
}

void printArray(int A[], int size) {
    for (auto i = 0; i < size; i++) {
        cout << A[i] << " ";
    }
}


// Oscillated Sorting
class Tape {
public:
    vector<vector<int>> vData;

    Tape() = default;

    void SetNewData(int data) {
        vector<int> new_data { data };
        this->vData.push_back(new_data);
    }

    vector<int> Pop() {
        vector<int> tmp;
        if (!vData.empty()) {
            tmp = vData.back();
            vData.pop_back();
        } else {
            tmp = vector<int>();
        }
        return tmp;
    }

    vector<int> GetData() {
        if (!vData.empty()) {
            return vData.back();
        } else {
            return {};
        }
    }

    void ChangeBack(const vector<int>& new_vector) {
        if(!vData.empty()) {
            vData.pop_back();
        }
        vData.push_back(new_vector);
    }

    friend ostream& operator<< (ostream& out,const Tape& tape) {
        for(const auto & i : tape.vData) {
            for(int j : i) {
                cout << j << " " << endl;
            }
            cout << endl;
        }

        return out;
    }
};

void PrintTapes(vector<Tape> tapes) {
    for(int i=0; i< tapes.size(); ++i) {
        cout << "TAPE " << i << endl;
        cout << tapes[i];
    }
}

vector<int> merge(vector<int> a, vector<int> b) {
    vector<int> result_vector;
    for(auto i = a.begin(); i != a.end(); i != a.end() ? ++i : i) {
        for(auto j = b.begin(); j != b.end() && i != a.end();) {
            if(*i < *j) {
                result_vector.push_back(*i);
                i = a.erase(i);
            } else {
                result_vector.push_back(*j);
                j = b.erase(j);
            }
        }
    }

    result_vector.insert(result_vector.end(), a.begin(), a.end());
    result_vector.insert(result_vector.end(), b.begin(), b.end());
    return result_vector;
}

int main() {
    cout << "------------ Balanced MergeSort ------------ (count: 20)\n";
    int arr[] = { 54, 26, 93, 17, 77, 31, 44, 55, 20, 40, 20, 55, 1, 7, 39, 11, 25, 19, 34, 10 };
    auto arr_size = sizeof(arr)/sizeof(arr[0]);

    cout << "Given array is \n";
    printArray(arr, arr_size);

    mergeSort(arr, 0, arr_size - 1);

    cout << "\nSorted array is \n";
    printArray(arr, arr_size);


    cout << "\n\n------------ Oscillated Sorting (MergeSort) ------------ (count: 20)\n";
    vector<Tape> tapes;

    vector<int> data;

    cout << "BEFORE SORTING: " << endl;
    for(int i = 0; i < 20; i++) {
        data.push_back(rand() % 1000);
        cout << data.back() << endl;
    }

    //add tapes
    tapes.reserve(NUM_OF_TAPES);
    for(int i = 0; i < NUM_OF_TAPES; ++i) {
        tapes.emplace_back();
    }

    int startTape = 0;
    while(!data.empty()) {
        //select data and write to tapes
        for(int i = startTape; i < startTape + NUM_OF_TAPES - 1 && !data.empty(); ++i) {
            int selected_tape = i % NUM_OF_TAPES;
            tapes[selected_tape].SetNewData(data.back());
            data.pop_back();

            cout << "Write data to tape #" << selected_tape << endl;
        }

        //merge data
        int tape_to_merge = (startTape + NUM_OF_TAPES - 1) % NUM_OF_TAPES;
        for(int i = startTape; i < startTape + NUM_OF_TAPES - 1; ++i) {
            int selected_tape = i % NUM_OF_TAPES;
            tapes[tape_to_merge].ChangeBack(merge(tapes[tape_to_merge].GetData(),
                                                  tapes[selected_tape].Pop()));

            cout << "Merge data from tape #" << selected_tape << " to tape #" << tape_to_merge << endl;
        }

        startTape++;
    }

    //final merge
    if(count_if(tapes.begin(),tapes.end(), [] (const Tape& tape) { return !tape.vData.empty(); }) > 1) {
        int tape_to_merge = (startTape + NUM_OF_TAPES - 1) % NUM_OF_TAPES;
        for(int i = startTape; i < startTape + NUM_OF_TAPES - 1; ++i) {
            int selected_tape = i % NUM_OF_TAPES;
            tapes[tape_to_merge].ChangeBack(merge(tapes[tape_to_merge].GetData(),tapes[selected_tape].Pop()));

            cout << "Merge data from tape #" << selected_tape << " to tape #" << tape_to_merge << endl;
        }
    }
    cout << "AFTER SORTING (TAPES STATE): " << endl;
    PrintTapes(tapes);

    return 0;
}