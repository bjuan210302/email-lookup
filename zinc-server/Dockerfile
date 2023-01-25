FROM busybox

WORKDIR "/zinc-bin"
COPY ./zinc .

ARG user=admin
ARG pass=Complexpass#123
ENV ZINC_FIRST_ADMIN_USER=$user
ENV ZINC_FIRST_ADMIN_PASSWORD=$pass

EXPOSE 4080

CMD ["./zinc"]