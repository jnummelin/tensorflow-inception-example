## Tensorflow models in containers

This repo has an example how to utilize TensorFlow models with containers.

## The model

The example inception model for image classification can be built as container image using `Dockerfile.model` file. It will download the model and put it to be served by tensorrflow-service component.

## API

The model servicing API is somewhat complex to use, so there's a simplified Go API implemented on top of it. Build it using `Dockerfile`.

## Deploying

There's a ready stack file to deploy everything on Kontena, so get you Kontena up-and-running. See https://kontena.io/docs/quick-start.html to get up-and-running in no time.

After you have Kontena platform running, you can install the stack with `kontena stack install jussi/tensorflow-example`

To test it out, you can use e.g. curl:
```
$ curl -s -XPOST -F "file=@/Users/jussi/Downloads/cropped_panda.jpg" image-classifier.kontena.works/classify | jq .
[
  {
    "Class": "giant panda, panda, panda bear, coon bear, Ailuropoda melanoleuca",
    "Score": 9.546637
  },
  {
    "Class": "indri, indris, Indri indri, Indri brevicaudatus",
    "Score": 6.6261067
  },
  {
    "Class": "gibbon, Hylobates lar",
    "Score": 4.3301826
  },
  {
    "Class": "lesser panda, red panda, panda, bear cat, cat bear, Ailurus fulgens",
    "Score": 4.0944114
  },
  {
    "Class": "titi, titi monkey",
    "Score": 2.8160584
  }
]
```