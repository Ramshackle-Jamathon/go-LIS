package LIS

import (
    "math"
)

func LIS (s []int) []int{  
    l := 0;
    P := make([]int, len(s));
    M := make([]int, len(s) + 1);
    for i := 0; i < len(s); i++ {
        lo := 1;
        hi := l;
        for lo <= hi {
            mid := int(math.Ceil(float64((lo+hi)/2)))
            if(s[M[mid]] < s[i]){
                lo = mid + 1
            } else {
                hi = mid-1
            }
        }
        newL := lo;

        P[i] = M[newL-1];
        M[newL] = i;

        if newL > l{
            l = newL
        }
    }
    S := make([]int, l);
    k := M[l]
    for i := l-1; i >= 0; i--{
        S[i] = s[k]
        k = P[k]
    } 
    return S
}