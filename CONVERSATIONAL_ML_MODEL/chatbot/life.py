from transformers import AutoTokenizer, AutoModelForCausalLM

tokenizer = AutoTokenizer.from_pretrained("PygmalionAI/pygmalion-6b")

model = AutoModelForCausalLM.from_pretrained("PygmalionAI/pygmalion-6b")

# Set the persona for the model
persona = """
Assistant is a highly intelligent language model trained to comply with user requests.
Assistant is always helpful and polite, and it is always willing to go the extra mile to help the user.
Assistant is also very knowledgeable, and it can answer any question that the user asks.
"""

model.set_persona(persona)

# Start interacting with the model
while True:
  prompt = input("User: ")
  output = model.generate(
    input_ids=tokenizer(prompt, return_tensors="pt").input_ids,
    max_length=100,
    do_sample=True,
    top_k=5,
    temperature=0.7,
  )

  print("Assistant: ", tokenizer.decode(output[0], skip_special_tokens=True))
