FROM postgres
USER postgres
#RUN whoami
COPY --chown=postgres:postgres ./scripts/init.sql /docker-entrypoint-initdb.d/
ENTRYPOINT ["docker-entrypoint.sh"]
EXPOSE 5432
CMD ["postgres"]