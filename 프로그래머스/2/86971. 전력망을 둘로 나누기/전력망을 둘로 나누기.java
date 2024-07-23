import java.util.*;
class Solution {
    static int[][] arr;
    public int solution(int n, int[][] wires) {
        int answer = n;
        arr = new int[n+1][n+1];
        int[] result = new int[n];
        // 인접 행렬 생성
        for(int i=0; i<wires.length; i++){
            int x = wires[i][0];
            int y = wires[i][1];
            arr[x][y] = 1;
            arr[y][x] = 1;
        }
        // 연결을 하나 끊고 bfs함수 호출 후 다시 연결
        for(int i=0; i<wires.length; i++){
            int x = wires[i][0];
            int y = wires[i][1];
            arr[x][y] = 0;
            arr[y][x] = 0;

            result[i] = bfs(n, x);
            
            arr[x][y] = 1;
            arr[y][x] = 1;
        }

        // count값들의 최소값
        for (int i=0; i<result.length; i++) {
            int cal = n-2*result[i];
            if (cal < 0) {
                cal = cal * (-1);
            }
            if (cal < answer) {
                answer = cal;
            }    
        }
        return answer;
    }
    // 너비 탐색
    public int bfs(int n, int start) {
        int[] visited= new int[n+1];
        int count=1;
        
        Queue<Integer> queue= new LinkedList<>();
        queue.offer(start);
        
        while(!queue.isEmpty()){
            int point= queue.poll();
            visited[point]= 1;
            
            for(int i=1; i<=n; i++){
                if(arr[point][i]==1 && visited[i] != 1) {
                    queue.offer(i);
                    count++;
                }
            }
        }
        return count;
    }
}


