class Solution {
    public double solution(int[] numbers) {
        double answer = 0;
            if(numbers.length%2==0){
            answer = (numbers[0] + numbers[numbers.length-1])/2F;
        }else{
            answer = numbers[(numbers.length-1)/2];
        }
        
        return answer;
    }
}