#include <unistd.h>
#include <limits.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

int main(int argc, char *argv[]) {
  char input[_POSIX_PATH_MAX];

  getcwd(input, sizeof(input));
  strcat(input, "/input");

  FILE* fp = fopen(input, "r");

  size_t bufsiz = BUFSIZ;
  char *buf = malloc(bufsiz);
  int c;

  for (int skip = 'A'; skip <= 'Z'; skip++) {
    size_t len = 0;
    buf[0] = '\0';

    while (c = fgetc(fp), isalpha(c)) {
      if (toupper(c) == skip) {
        continue;
      }

      if (len > 0) {
        int prev = buf[len - 1];
        if (prev != c && tolower(prev) == tolower(c)) {
          len--;
          continue;
        }
      }

      if (len+1 > bufsiz) {
        bufsiz += BUFSIZ;
        buf = realloc(buf, bufsiz);
      }

      buf[len++] = c;
    }

    printf(
      "After removing unit type '%c', and after possible reactions, the resulting polymer contains %lu units.\n",
      skip,
      len
    );

    rewind(fp);
  }

  free(buf);

  fclose(fp);

  return 0;
}
