#include <stdio.h>
#include <stdlib.h>

#define ARRAY_SIZE(arr) (sizeof(arr) / sizeof((arr)[0]))

int
main(void)
{
    FILE *fp = fopen("test.txt", "w+b");
    if (!fp) {
        perror("fopen");
        return EXIT_FAILURE;
    }

    char cs[] = "Hello, world!\n";
    fwrite(cs, ARRAY_SIZE(cs), sizeof(*cs), fp); 

    unsigned char buffer[20];

    fseek(fp, 0, SEEK_SET);

    size_t ret = fread(buffer, sizeof(*buffer), ARRAY_SIZE(buffer), fp);
    if (ret <= 0) {
        perror("fread");
        fprintf(stderr, "fread() failed: %zu\n", ret);
        exit(EXIT_FAILURE);
    }

    printf("%s", buffer);

    fclose(fp);

    exit(EXIT_SUCCESS);
}