package A2;


public class A2 {
    public static void main(String[] args) {
        int x = 3;
        {
            System.out.println(x);
            x = 4;
            System.out.println(x);
        }
        System.out.println(x);
    }
}
