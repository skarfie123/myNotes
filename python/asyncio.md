# asyncio

[import asyncio](https://www.youtube.com/playlist?list=PLhNSoGM2ik6SIkVGXWBwerucXjgP1rHmB)

## event loop

get the current loop, or create if there isn't one

```python
loop = asyncio.get_event_loop()
```

 run the loop for ever

```python
loop.run_forever()
```

run the loop for 5 seconds

```python
loop.run_until_complete(asyncio.sleep(5))
# or
loop.call_later(5, loop.stop)
loop.run_forever()
```

call a function soon or after some time. these only run if when the loop is running.

```python
loop.call_soon(lambda: print("hello asap"))
loop.call_later(1, lambda: print("hello at least 1s later"))
loop.run_until_complete(asyncio.sleep(5))
```

use debug mode when testing. this will warn about any tasks that cause event loop starvation (hogging). alternatively see [here](https://docs.python.org/3/library/asyncio-dev.html#debug-mode)

```python
loop.set_debug(True)
```

in production use 3rd party `uvloop`, which is much more performant

```python
import uvloop
uvloop.install() # must be run before loop is created
```

## Coroutines

run an async function

```python
asyncio.run(func())
```

run with a 10s timeout

```python
asyncio.run(asyncio.wait_for(func(), 10))
```

use a main function to catch timeouts gracefully

```python
async def async_main() -> None:
    try:
        await asyncio.wait_for(func(), 10)
    except asyncio.TimeoutError:
        ...

asyncio.run(async_main())
```

await multiple coroutines at the same time.

```python
asyncio.gather( )
```

### Type hierarchy

- collections.abc.Awaitable
  - collections.abc.Coroutine
  - asyncio.Future
    - asyncio.Task

An async function is a callable that returns a coroutine, which is an awaitable. A coroutine can only be awaited once.

### Debugging

In debug mode (eg. with `PYTHONASYNCIODEBUG=1`) asyncio will warn if coroutines are not awaited. The warning will be more detailed with `PYTHONTRACEMALLOC=1`.

## Futures

Futures are awaitable objects that represent a future value.

Set value

```python
fut.set_result(value)
```

Set exception

```python
fut.set_exception(exc)
```

Cancel future

```python
fut.cancel()
```

Get result. this either:

- returns set value
- raises set exception
- raises CancelledError
- raises InvalidStateError if not yet awaitable

```python
fut.result()
```

add a callback (`Callable[[Future], Any]`) to be run async when the future is done

```python
fut.add_done_callback(callback)
```

## Tasks

Coroutines is a builtin type that is not aware of asyncio concepts nor have extra features such as:

- the event loop
- exceptions raised
- cancellations (a special exception)
- results returned

create a background task:
this task will be excuted whenever we are waiting for something else, eg. a sleep.

```python
task = asyncio.create_task(func())
```

wait 5s for some tasks to finish:

```python
done, pending = await asyncio.wait(tasks, timeout=5)
```
