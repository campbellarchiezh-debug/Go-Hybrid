package main

// W-024: Disjoint Parallel Slices via Capacity Clipping
fn parallel_sum(s []int) int {
    if len(s) <= 100 {
        mut sum := 0
        for x in s {
            sum += x
        }
        return sum
    }
    mid := len(s) / 2
    // Lowered to 3-index full slice expressions
    split s at mid into left, right {
        task sum1 := parallel_sum(move left)
        task sum2 := parallel_sum(move right)
        return sum1 + sum2
    }
}
