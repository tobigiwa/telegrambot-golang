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


from transformers import AutoTokenizer, AutoModelForCausalLM
import torch

tokenizer = AutoTokenizer.from_pretrained("microsoft/DialoGPT-medium")
model = AutoModelForCausalLM.from_pretrained("microsoft/DialoGPT-medium")

# Load conversation history from JSON file
with open("conversation_history.json", "r") as f:
    conversation_history = json.load(f)

# Fine-tune the model on the conversation history
for i in range(len(conversation_history)):
    input_text = conversation_history[i]["user_message"]
    output_text = conversation_history[i]["model_response"]

    # Encode the input and output text
    input_ids = tokenizer.encode(input_text + model.config.eos_token, return_tensors="pt")
    output_ids = tokenizer.encode(output_text + model.config.eos_token, return_tensors="pt")

    # Concatenate the input and output text
    input_ids = torch.cat([input_ids, output_ids], dim=-1)

    # Train the model on the concatenated text
    loss = model(input_ids=input_ids, labels=input_ids).loss

    # Print the loss for debugging purposes
    print(f"Loss at step {i}: {loss.item()}")

# Save the fine-tuned model
model.save_pretrained("fine_tuned_model")



from transformers import AutoTokenizer, AutoModelForCausalLM, AdamW
import torch

tokenizer = AutoTokenizer.from_pretrained("microsoft/DialoGPT-medium")
model = AutoModelForCausalLM.from_pretrained("microsoft/DialoGPT-medium")

# Load conversation history from JSON file
with open("conversation_history.json", "r") as f:
    conversation_history = json.load(f)

# Set up the optimizer
optimizer = AdamW(model.parameters(), lr=1e-4)

# Fine-tune the model on the conversation history
for i in range(len(conversation_history)):
    input_text = conversation_history[i]["user_message"]
    output_text = conversation_history[i]["model_response"]

    # Encode the input and output text
    input_ids = tokenizer.encode(input_text + model.config.eos_token, return_tensors="pt")
    output_ids = tokenizer.encode(output_text + model.config.eos_token, return_tensors="pt")

    # Concatenate the input and output text
    input_ids = torch.cat([input_ids, output_ids], dim=-1)

    # Train the model on the concatenated text
    loss = model(input_ids=input_ids, labels=input_ids).loss

    # Backpropagate the loss and update the model parameters
    loss.backward()
    optimizer.step()
    optimizer.zero_grad()

    # Print the loss for debugging purposes
    print(f"Loss at step {i}: {loss.item()}")

# Save the fine-tuned model
model.save_pretrained("fine_tuned_model")




import tensorflow as tf
import transformers

# Load the model
model = transformers.AutoModelForCausalLM.from_pretrained("microsoft/DialoGPT-medium")

# Load the conversation history from JSON file
with open("conversation_history.json", "r") as f:
    conversation_history = json.load(f)

# Generate a response based on the conversation history and user prompt
input_text = ""
for i in range(len(conversation_history)):
    input_text += conversation_history[i]["user_message"] + model.config.eos_token
    input_text += conversation_history[i]["model_response"] + model.config.eos_token

user_prompt = input("Enter your prompt: ")
input_text += user_prompt + model.config.eos_token

input_ids = tf.convert_to_tensor(input_text)
response_ids = model.generate(input_ids, max_length=128, do_sample=True, top_k=5)

response = transformers.GPT2Tokenizer.from_pretrained("microsoft/DialoGPT-medium").decode(response_ids, skip_special_tokens=True)

# Print the response
print(response)


from transformers import AutoTokenizer, AutoModelForCausalLM
import torch
import json
tokenizer = AutoTokenizer.from_pretrained("microsoft/DialoGPT-medium")
model = AutoModelForCausalLM.from_pretrained("microsoft/DialoGPT-medium")
# Load conversation history from JSON file
with open("conversation_history.json", "r") as f:
conversation_history = json.load(f)
# Generate a response based on the conversation history and user prompt
input_text = ""
for i in range(len(conversation_history)):
input_text += conversation_history[i]["user_message"] + model.config.eos_token
input_text += conversation_history[i]["model_response"] + model.config.eos_token
user_prompt = input("Enter your prompt: ")
input_text += user_prompt + model.config.eos_token
input_ids = tokenizer.encode(input_text, return_tensors="pt")
output_ids = model.generate(input_ids)
output_text = tokenizer.decode(output_ids[0], skip_special_tokens=True)
print(output_text)



from transformers import AutoModelForCausalLM, AutoTokenizer
import torch


tokenizer = AutoTokenizer.from_pretrained("microsoft/DialoGPT-medium")
model = AutoModelForCausalLM.from_pretrained("microsoft/DialoGPT-medium")

# Let's chat for 5 lines
for step in range(5):
    # encode the new user input, add the eos_token and return a tensor in Pytorch
    new_user_input_ids = tokenizer.encode(input(">> User:") + tokenizer.eos_token, return_tensors='pt')

    # append the new user input tokens to the chat history
    bot_input_ids = torch.cat([chat_history_ids, new_user_input_ids], dim=-1) if step > 0 else new_user_input_ids

    # generated a response while limiting the total chat history to 1000 tokens, 
    chat_history_ids = model.generate(bot_input_ids, max_length=1000, pad_token_id=tokenizer.eos_token_id)

    # pretty print last ouput tokens from bot
    print("DialoGPT: {}".format(tokenizer.decode(chat_history_ids[:, bot_input_ids.shape[-1]:][0], skip_special_tokens=True)))


from transformers import AutoTokenizer, AutoModelForCausalLM
import torch
import json

tokenizer = AutoTokenizer.from_pretrained("microsoft/DialoGPT-medium")
model = AutoModelForCausalLM.from_pretrained("microsoft/DialoGPT-medium")

input_text: str = ""
for i in data:
  chat_history_ids += i["User"] + tokenizer.eos_token
  # input_text += i["Bot"] + tokenizer.eos_token

flag: bool = True
while  flag:
  user_prompt = input("Enter your prompt: ")
  new_user_input_ids = tokenizer.encode(user_prompt + tokenizer.eos_token, return_tensors='pt')

  bot_input_ids = torch.cat([chat_history_ids, new_user_input_ids], dim=-1)

  output_ids = model.generate(bot_input_ids, max_length=1000, pad_token_id=tokenizer.eos_token_id)

  output_text = tokenizer.decode(output_ids[:, bot_input_ids.shape[-1]:][0], skip_special_tokens=True)
  print("\n\n")
  print(output_text)
  print("\n\n")
  if user_prompt == "end":
    flag = False


  
  from transformers import AutoTokenizer, AutoModelForCausalLM
import torch
import json

tokenizer = AutoTokenizer.from_pretrained("microsoft/DialoGPT-medium")
model = AutoModelForCausalLM.from_pretrained("microsoft/DialoGPT-medium")

# Load conversation history from JSON file
# with open("ml-data.json", "r") as f:
#   conversation_history = json.load(f)


# Generate a response based on the conversation history and user prompt
input_text: str = ""
for i in data:
  input_text += i["user_message"] + tokenizer.eos_token
  input_text += i["model_response"] + tokenizer.eos_token


flag: bool = True

while  flag:
  user_prompt = input("Enter your prompt: ")
  input_text += user_prompt + tokenizer.eos_token


  input_ids = tokenizer(input_text, return_tensors="pt")["input_ids"]

  output_ids = model.generate(input_ids)

  output_text = tokenizer.decode(output_ids[0], skip_special_tokens=True)
  print("\n\n")
  print(output_text)
  print("\n\n")
  if user_prompt == "end":
    flag = False
