package A3;

public class Main {
    public static void main(String[] args) {
        for (int i = 0; i <= 1025; i++) {
            if (!hasExactFloatRepresentation(i)) {
                System.out.print(i + ", ");
            }
        }
    }

    //Hinweis: Die Zahlen 0 und 1 haben jeweils exakte Darstellungen als float.
    private static boolean hasExactFloatRepresentation(int n) {
        if (n == 0 || n == 1) {
            return true;
        }
        int exponent = getExponent(n);
        if (exponent >= 23) {
            return true;
        }
        int powerOfTwo = 1 << exponent;
        return powerOfTwo == n || powerOfTwo == n + 1;
    }

    private static int getExponent(int n) {
        if (n == 0) {
            return 0;
        }
        int exponent = 1;
        while ((n & (1 << exponent)) == 0) {
            exponent++;
        }
        return exponent;
    }
}


