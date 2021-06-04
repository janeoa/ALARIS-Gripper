#include <iostream>
#include <math.h>

#include <cstdlib>
#include <ctime>

using namespace std;


bool is_next_on_right(int prev, int curr){
	bool is_right = false;
	
	int delta = curr - prev;
	if(delta > 0 && delta < 4) is_right = true;
//	if(delta < 0 && delta > -4) is_right = false;
	
	if(delta < -4 && delta > -8) is_right = true;
//	if(delta < 0 && delta > -4) is_right = false;
	
	return is_right;
}


void test(int prev, int curr, bool should){
	string str = (is_next_on_right(prev, curr))?"right":"left ";
	string st2 = (should)?"right":"left ";
	
	string st3 = (is_next_on_right(prev, curr) == should)?"✅":"❌";
	
	cout << prev << "->" << curr << " is [" << str << "], should be ["<< st2 << "]   "<< st3 <<  endl;
}



int main(int argc, char *argv[]) {
	test(4,9,0);
	test(9,4,1);
	test(7,8,1);
	test(6,5,0);
	test(2,5,1);
	test(7,0,1);
	
	
}


