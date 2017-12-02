#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#define bins 8
#define bit_size 32
unsigned int bit_vector[bins] = {};

int set_bit(unsigned int *bit_vector, int pos)
{
        if (pos < 0 || pos > 255) return -1;
        int offset = pos/bit_size;
        int shift = pos%bit_size;
        bit_vector[offset] |= 1 << shift;
        return 1;
}

int is_set(unsigned int *bit_vector, int pos)
{
    if (pos < 0 || pos > 255) return -1;
    int offset = pos/bit_size;
    int shift = pos%bit_size;
    return ((bit_vector[offset] & (1<< shift))?1:0);
}

char* removeDuplicateLetters(char* s) {
    int i, j, k;
    char *str = s;
    i = j = 0;
    while (str[i] != '\0') {
        if (is_set(bit_vector, (int)str[i]))  {
            i++;
        } else {
            set_bit(bit_vector, (int)str[i]);
            str[j++] = str[i++];
        }
    }
    for (j=0,k=0; k<256;k++) {
        if (is_set(bit_vector, k)) { 
            str[j++] = k;
        }  
    }
    while (j<i) str[j++] = 0;
    return s;
}

int main()
{
    int res; 
    char str[10];
    printf("Enter the string (Extended ASCII 255) :");
    scanf("%9s", str);
    
    // Findind duplicates in a string
    /*
    res = find_duplicates(str);
    if (res) {
        printf("Duplicates present in %s", str);
        return 1;
    } else {
        printf("Unique string");
        return 1;
    }*/
    
    removeDuplicateLetters(str);
    printf("\n final: %s", str);
    
   /* res = set_bit(bit_vector, 0); 
    printf("255: %d\n", is_set(bit_vector, 0));
    res = set_bit(bit_vector, 1); 
    printf("32: %d\n", is_set(bit_vector, 1));
    res = set_bit(bit_vector, 2); 
    printf("31: %d\n", is_set(bit_vector, 2));
    res = set_bit(bit_vector, 3); 
    printf("24: %d\n", is_set(bit_vector, 3));
    res = set_bit(bit_vector, 4); 
    printf("0: %d\n %d\n", is_set(bit_vector, 4), *(unsigned int *)bit_vector);
    res = clear(bit_vector, 4);
    printf("0: %d\n %d", is_set(bit_vector, 4), *(unsigned int *)bit_vector);
    */
    return 1;
}
