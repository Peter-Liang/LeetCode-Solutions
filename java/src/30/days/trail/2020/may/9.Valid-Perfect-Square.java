class Solution {
    public boolean isPerfectSquare(int num) {
        if (num <= 1) return true;

        int left = 0;
        int right = num / 2;
        while (left <= right) {
            int mid = (left + right) / 2;
            if (mid == 0) return false;
            int div = num / mid;

            if (div * div == num) return true;

            if (div < mid) {
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }

        return false;
    }
}
