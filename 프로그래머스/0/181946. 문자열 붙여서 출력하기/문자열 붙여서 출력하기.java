import java.util.Scanner;

// 연산자
public class Solution1 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String a = sc.next();
        String b = sc.next();
        
        System.out.println(a+b);
    }
}
// StringBuilder
public class Solution2 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String a = sc.next();
        String b = sc.next();
        
        StringBuilder sb = new StringBuilder(10);
        sb.append(a);
        sb.append(b);
        
        System.out.println(sb.toString());
    }
}

// String.concat()
public class Solution3 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String a = sc.next();
        String b = sc.next();
        
        String str = a.concat(b);
        System.out.println(str);
    }
}

// String.format()
public class Solution4 {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String a = sc.next();
        String b = sc.next();
        
        String str = String.format("%s%s", a, b);
        System.out.println(str);
    }
}