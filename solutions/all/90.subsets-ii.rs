impl Solution {
    pub fn subsets_with_dup(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut sorted_nums = nums.to_vec();
        sorted_nums.sort();
        let mut result = Vec::new();

        get_subsets(&mut result, &sorted_nums, Vec::new(), 0);
        return result;
    }
}

fn get_subsets(result: &mut Vec<Vec<i32>>, nums: &Vec<i32>, cur: Vec<i32>, start: usize) {
    result.push(cur.clone());

    for i in start..nums.len() {
        if i == start || nums[i] != nums[i - 1] {
            let mut vec = cur.clone();
            vec.push(nums[i]);
            get_subsets(result, &nums, vec, i + 1);
        }
    }
}
