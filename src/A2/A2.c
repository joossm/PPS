#include <stdio.h>
int main (int argc, char **argv) {
  int x = 3;
  { printf ("%d\n", x);
    int x = 4;
    printf ("%d\n", x);
  }
  printf ("%d\n", x);
}

