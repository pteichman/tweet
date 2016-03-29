tweet is a command line program for posting to Twitter.

To install:

    $ go get github.com/pteichman/tweet

To use:

    $ tweet -c creds.conf "Had pancakes for breakfast #winning"

The config file contains your Twitter credentials. It's up to you to go through the OAuth dance to get them. It looks like:

    {
        "ConsumerKey": "....",
        "ConsumerSecret": "....",
        "AccessKey": "....",
        "AccessSecret": "...."
    }
