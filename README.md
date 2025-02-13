# hnews
Download HackerNews Post as JSON

## Installation
To install the tool locally, run:
```bash
git clone https://github.com/user/hnews.git
cd hnews
go build -o hnews .
```

## Usage
To download a Hacker News post, run:
```bash
./hnews get --id $POST_ID
```
Replace `$POST_ID` with the ID of the Hacker News post you want to download. It will output the post details in JSON format.

## Example
For example, to download item #8863, use:
```bash
% ./hnews get --id 8863
{
  "text": "It's a deal!"
}
```