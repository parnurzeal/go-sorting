#include<iostream>
#include<vector>


using namespace std;

int main(void){
  vector<int> nums;
  int tmpIn;
  while(cin>>tmpIn){
    nums.push_back(tmpIn);
  }
  for(int i = 0;i<nums.size();i++){
    cout<<nums[i]<<" ";
  }




  return 0;
}
