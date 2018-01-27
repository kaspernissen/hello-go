FROM scratch

WORKDIR /hello
COPY hello ./hello

CMD ["./hello"]
