# he challenge is to create a function that determines whether a given string is a palindrome or not.
# A palindrome is a string that reads the same forwards and backwards, ignoring spaces, punctuation, and capitalization.
# Requirments:
#     - Ignore spaces, punctuation, and capitalization.
#     - You can assume the input string will only contain alphanumeric characters and spaces.


def is_palindrome(s: str) -> bool:
    # Remove non-alphanumeric characters and convert to lowercase
    cleaned = ''.join(char.lower() for char in s if char.isalnum())
    # Compare the cleaned string with its reverse
    print(cleaned)
    print(cleaned[::-1])
    return cleaned == cleaned[::-1]

is_palindrome("A man a plan a canal Panama")