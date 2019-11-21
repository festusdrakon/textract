import en_core_web_sm


nlp = en_core_web_sm.load()
with open('./fields/four1.jpg_data.txt', 'r') as fp:
    for line in enumerate(fp):
        line = fp.readline()
        doc = nlp(line)
#        print([(w.text, w.pos_) for w in doc])
        for w in doc:
            if w.pos_ == 'NUM':
                print(w)
