FROM busybox

COPY ./zinc /zinc-bin

ENV ZINC_FIRST_ADMIN_USER=admin
ENV ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123

EXPOSE 4080

WORKDIR "/zinc-bin"
CMD ["./zinc"]