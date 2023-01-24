FROM busybox

COPY ./zinc /zinc-bin

ARG user=admin
ARG pass=Complexpass#123
ENV ZINC_FIRST_ADMIN_USER=$user
ENV ZINC_FIRST_ADMIN_PASSWORD=$pass

EXPOSE 4080

WORKDIR "/zinc-bin"
CMD ["./zinc"]