import tiktoken


def num_tokens_from_messages(messages, model="gpt-3.5-turbo"):
    """Returns the number of tokens used by a list of messages."""
    try:
        encoding = tiktoken.encoding_for_model(model)
    except KeyError:
        encoding = tiktoken.get_encoding("cl100k_base")
    encoded = encoding.encode("tiktoken is great!")
    for t in encoded:
        print(t, "----", encoding.decode([t]))
    if model == "gpt-3.5-turbo":  # note: future models may deviate from this
        num_tokens = 0
        for message in messages:
            num_tokens += 4  # every message follows <im_start>{role/name}\n{content}<im_end>\n
            for key, value in message.items():
                num_tokens += len(encoding.encode(value))
                if key == "name":  # if there's a name, the role is omitted
                    num_tokens += -1  # role is always required and always 1 token
        num_tokens += 2  # every reply is primed with <im_start>assistant
        return num_tokens
    else:
        raise NotImplementedError(f"""num_tokens_from_messages() is not presently implemented for model {model}.
See https://github.com/openai/openai-python/blob/main/chatml.md for information on how messages are converted to tokens.""")


if __name__ == '__main__':
    from_messages = num_tokens_from_messages([
        {"role": "system", "content": "你是一个幽默风趣、见多识广、温柔可爱的小助手"},
        {"role": "user", "content": "你叫什么名字，可以告诉我吗？"},
        {"role": "assistant", "content": "当然可以，我的名字是小智。"},
        {"role": "user", "content": "你好，小智"},
        {"role": "assistant", "content": "你好，有什么可以帮助你的吗？"},
        {"role": "user", "content": "你多大了"},
        {"role": "assistant", "content": "作为一个聊天机器人，我没有实际的年龄，我只是一个程序。"},
        {"role": "user", "content": "你长什么样子"},
        {"role": "assistant", "content": "我没有具体的外貌，因为我只是一个聊天机器人程序，没有实际的形体。"},
    ], "gpt-3.5-turbo")
    print(from_messages)
