import asyncio


async def keep_printing(name="default"):
    while True:
        print(name)
        await asyncio.sleep(0.5)


async def async_main():
    try:
        await asyncio.wait_for(
            asyncio.gather(
                keep_printing(1),
                keep_printing(2),
                keep_printing(3),
            ),
            timeout=3,
        )
    except asyncio.TimeoutError:
        print("Timeout!")


def main():
    asyncio.run(async_main(), debug=True)


if __name__ == "__main__":
    main()
