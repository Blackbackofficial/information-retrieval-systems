#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

static int NUM_OF_TAPES = 4;

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