# jackalope
A relational json database for all kinds of state

> [!WARNING]  
> This framework is under heavy development and is not ready for production

I was looking at building an ORM or using GOORM initially. But The following are reasons why building my own database is advantageous:
- SQLite can only be used with Cgo
- Building something like litesail is easier if I don't use SQLite
- I can also better address the problem of saving uploaded files to the database. (Images, sound, ect.)
- We can have native type support. One example is that SQLite does not support arrays
- I care less about this one but it's fun to think that this approach might be faster than the convention
- Building a database in this way also gives us the freedom to implement things like ques. All of the state can be contained here for ease of use

I tried building this in binary first but I ran into problems with vars needing to be fixed length (So I had to use something like [50]byte instead of a string. I don't like it)
So I'm thinking that the best route will be storing data in json and other files in a folder.
This has the side benefit of data being recoverable. And uploading to s3 is like a regular backup

Also, it would be nice to build this in a way that there could be a Python implementation in the future. I could have used this with my previous Flask projects

Overview:
Define "tables" with structs (or Classes in Python's case)
Field constraints/details. unique, default values, many to many
Common dynamic types for fields like String, Integer, Float, list, Image/Bin (All mutable)
An interface for creating, reading, updating, and deleting records

Build indexes on startup and make sure all records and files are intact. (The user should not edit the database files directly)


Later: https://sqldocs.org/sqlite/sqlite-write-ahead-logging/
