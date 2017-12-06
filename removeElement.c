/*
*
* Vasu Mahalingam, 2017
*
*
*
*Given an array and a value, remove all instances of that value in-place and return the new length.
*Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
*The order of elements can be changed. It doesn't matter what you leave beyond the new length
*/

int removeElement(int* nums, int numsSize, int val) {
    int i = 0, j = 0;
    for (;i < numsSize;) {
        if (nums[i] != val) {
                nums[j++] = nums[i++]; 
        } else {
            i++;
        }
    }
    return j;
}
