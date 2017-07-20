 ======
 README 
 ======

 Version 1.2, March 2010

 . What is this?

    This dataset contains <user, artist, plays> tuples (for ~360,000 users) collected from Last.fm API,
    using the user.getTopArtists() method.

 . Files:

    usersha1-artmbid-artname-plays.tsv (MD5: be672526eb7c69495c27ad27803148f1)
    usersha1-profile.tsv               (MD5: 51159d4edf6a92cb96f87768aa2be678)
    mbox_sha1sum.py                    (MD5: feb3485eace85f3ba62e324839e6ab39)

 . Data Statistics:

    File usersha1-artmbid-artname-plays.tsv:

      Total Lines:           17,559,530
      Unique Users:             359,347
      Artists with MBID:        186,642
      Artists without MBID:     107,373

 . Data Format:

    The data is formatted one entry per line as follows (tab separated "\t"):

    File usersha1-artmbid-artname-plays.tsv:
      user-mboxsha1 \t musicbrainz-artist-id \t artist-name \t plays

    File usersha1-profile.tsv:
      user-mboxsha1 \t gender (m|f|empty) \t age (int|empty) \t country (str|empty) \t signup (date|empty)

 . Example:

    usersha1-artmbid-artname-plays.tsv:
      000063d3fe1cf2ba248b9e3c3f0334845a27a6be \t a3cb23fc-acd3-4ce0-8f36-1e5aa6a18432 \t u2 \t 31
      ...

    usersha1-profile.tsv
      000063d3fe1cf2ba248b9e3c3f0334845a27a6be \t m \t 19 \t Mexico \t Apr 28, 2008
      ...

 . License:

    The data contained in lastfm-dataset-360K.tar.gz is distributed with permission of Last.fm. 
    The data is made available for non-commercial use.
