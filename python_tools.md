# Python Tools

## Debugging

### Logaru

```python
from loguru import logger

@logger.catch # Add this to track errors
def main():
    # something that makes errors
```

![logaru](https://miro.medium.com/max/638/1*Wa4aTPXFq5Zv5bAKRVt3AA.png)

### Snoop

```python
import snoop 

@snoop
def function(x):
    # some complex function
```

![snoop](https://miro.medium.com/max/655/1*SREOQFmPpfZQ6IGB-ozVAQ.png)

### Heartrate

```python
import heartrate 
heartrate.trace(browser=True)
```

![heartrate](https://miro.medium.com/max/700/1*iRPyzFsJ_kz7dcZ9Y7QVuQ.png)

### icecream

```python

from icecream import ic 

ic(num)
ic(func(num))
ic()
def time_format():
    return f'{datetime.now()}|> '
ic.configureOutput(prefix=time_format)
ic.configureOutput(includeContext=True)
```

![icecream](https://miro.medium.com/max/700/1*ZLVC2GFBc54M3WNtOfvWoA.png)

## Profiling

### CPython Built-In

A whole script: `python -m cProfile -o blah.prof blah.py`

or just a portion of code, e.g. the `blah` function:

```python
import cProfile
import pstats

with cProfile.Profile() as pr:
    blah()

stats = pstats.Stats(pr)
stats.sort_stats(pstats.SortKey.TIME)
# stats.print_stats()
stats.dump_stats(filename="blah.prof")
```

### Snakeviz

`snakeviz blah.prof`

![snakeviz](https://jiffyclub.github.io/snakeviz/img/icicle.png)

## Tools

### HTTP Server

`py -m http.server`

### ptipython

`ptipython`

![ptipython](https://i.imgur.com/OJMYYZY.png)
