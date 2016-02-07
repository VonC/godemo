FROM scratch

COPY helloworldwebhttps .
COPY localhost.key .
COPY localhost.crt .

EXPOSE 80
EXPOSE 443

ENTRYPOINT ["/helloworldwebhttps"]
CMD ["80", "443"]
