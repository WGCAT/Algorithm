import java.util.*;

class Solution {
    public int solution(int[][] maps) {
        int answer = 0;
        int n = maps.length;
        int m = maps[0].length;
        
        // 이동 가능한 방향 (상,하,우,좌)
        int[][] directions = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        
        Queue<int[]> queue = new LinkedList<>();
        // 현재위치에 시작 위치(0, 0)를 큐에 추가 후 방문처리
        queue.offer(new int[]{0, 0, 1}); 
        // 큐가 비어있지 않은 동안 내부에서 계속 노드를 생성하여 큐에 넣는다
        while (!queue.isEmpty()) {
            // 현재위치 반환
            int[] currentPosition = queue.poll();
            int x = currentPosition[0];
            int y = currentPosition[1];
            int distance = currentPosition[2];

       	    // 목적지에 도달하면 최단 거리를 반환후 종료
            if (x == n - 1 && y == m - 1) { 
                answer = distance;
                break;
            }

            // 좌표 탐색
            for (int[] dir : directions) {
                int newX = x + dir[0];
                int newY = y + dir[1];

                // 이동 가능한 위치이면 큐에 추가 및 방문처리 후 벽 생성
                if (newX >= 0 && newX < n && newY >= 0 && newY < m && maps[newX][newY] == 1) {
                    queue.offer(new int[]{newX, newY, distance + 1});
                    maps[newX][newY] = 0;
                }
            }
        }
        // 목적지에 도달할 수 없는 경우 -1 반환
        if (answer == 0) {
            return -1;
        }
        
        return answer;
    }
}