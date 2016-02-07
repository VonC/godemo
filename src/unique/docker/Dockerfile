FROM scratch

COPY helloworldwebhttps .
COPY localhost.key .
COPY localhost.crt .

ENTRYPOINT ["/helloworldwebhttps"]
CMD ["80", "443"]
