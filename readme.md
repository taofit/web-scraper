**Web Scraper in GoLang**

This little program uses Colly library to implement a web scraper.

It will go to a web page and fetch each row in a table element and store the row data in a csv file.


## Steps to run the program
1) Create a dockerfile under the root folder which is the parent folder of emitter and countd folders. The program contains a dockerfile which can build an image that contains two services that are emitter and countd. These two services will talk to each other via 0mq.
2) Build image under root folder where dockerfile is, enter command: `docker build --platform amd64 --tag 0mq-golang-assignment .`, `0mq-golang-assignment` is the image name. On Mac M1 it requires to enter `--plateform amd64` to build a docker image, other computer may not need this flag.
3) Run container from the image in interactive mode in the current terminal, enter command: `docker run -it 0mq-golang-assignment`. and now we are in docker container environment, you should see `/usr/src/app/emitter#` in the terminal, enter command: `cd /countd`, it is to project's countd folder, then enter command: `ls` to list all the files under root. enter command: `./0MQ-golang-countd`, it is to open the countd service, which is listening the message sending from another terminal that runs emitter service.
4) Open a new terminal, enter command: `docker ps` and see something like:
`CONTAINER ID   IMAGE     COMMAND           CREATED          STATUS    PORTS         NAMES <br />  
6a61bfd343a8   0mq-golang-assignment      "bash"           16 seconds ago   Up 16  seconds                                                     mystifying_payne`.
We can see the running container id `6a61bfd343a8` from the above result.

5) In this terminal, enter command `docker exec -it 6a61bfd343a8 bash`
6) Now we have the same container opens in another terminal, we will open emitter in this terminal, the emitter will send message to countd that is already open in the first terminal.
7) In this termival, you should see `/usr/src/app/emitter#` and run command: `./0MQ-golang-emitter`, then you can enter the text and start sending message to countd service opened in another terminal.

## Solution
Code for aggregates is in countd/aggragation folder, 
If a new word comes in, the program will do flushing immediately without inserting the new word information to the word storage. Once both the flushing and starting another go routine for the (1s or 10s) intervals flushing are done, the new word information will be saved to word storage.

If an existing word comes in, the program will simply increase the both counts by 1

Since there are so many flushes triggerd by different go routines, and they can happen simultaneously, so reading and writing to the map must be locked, it can achieved by using sync.RWMutex.

The receiver is rather simple, and is done by calling printWords() of the storage struct.
Receiver will print the timestamp for first and last seen, which is based on the aggregates format, but it could print human readable format which can be achieved by `fmt.Println(current_time.Format("2022-10-02 12:07:23"))`, 'current_time' should be 'time.Now()'. It may not need to be put in wordElement struct, so for simplicity purpos, just skip it.
