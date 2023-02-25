**Manual Way:**

      a. Install PostgreSQL on Host PC.
   
      b. Install Golang on Host PC.

     Then:

      1. Build Project by using command `go build -v .`
      2. Configure DB by changing `config.yml` file in `local_database` section. 
      3. Change `current_db`'s `RUNNING` value from docker_database to `local_database`.
      4. Run the project by using command `./event_bright serve`.