//NOTE:
/\*
Routines -
syntax: go func()

    -go routine size-		2/4kb - as per the program requires,
    - OS threads size - 1MB - ,,
    -the order in which the routines execute can't be determined.
    -they are concurrent(software based),no parallel(hardware based)
    -By default go runs as a main go routine,and the other routines created run concurrently with it,
    	when the main go routine finishes it execution ,all the other routines are killed.

Channels -
syntax: ch:=make(chan data_type, buffer)
go func1(arg , ch)
go func2(arg , ch)
x <-ch
y <-ch
-Used to create a communication between 2 routines . - Other languages emphasize on shared memory between threads,But go emphasize on communication between routines.
-Buffered vs Unbuffered channels
-buferred => can have a buffer of some mentioned values stored inside before pushing again.(asynchronous)
-Unbuffered => to keep the routines synchronised (i.e) pushes only after it is received(synchronous).

Mutex - //NOTE:Locks dont lock the variables,its just the mental model , the other thread doesnt access the variable untill it ses the mu.unlock ,so that it could mu.lock ()
EG: think of a single kay that users have to use to access a resource, if one user is using the key, the other user has to wait until the first user finishes and releases the key.
same as in OS but for routines.
syntax : mu sync.Mutex

\*/

Implent this:
-single pathway
-alice& bob (def n blind)
both go left n right at the same time (innfinite loop),
