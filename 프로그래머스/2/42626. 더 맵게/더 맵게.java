import java.util.*;

class Solution {
    public int solution(int[] scoville, int K) {
        int answer = 0;
        PriorityQueue<Integer> queue = new PriorityQueue<>();
        for (Integer s : scoville) {
            queue.add(s);
        }
        while (queue.peek() < K) {
            if (queue.size()==1) {
                return -1;
            }else{
                queue.add(queue.poll()+queue.poll()*2);
                answer++;
            }
        }
        
        return answer;
    }
}