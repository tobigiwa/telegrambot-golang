from transformers import AutoTokenizer, AutoModelForSeq2SeqLM

tokenizer = AutoTokenizer.from_pretrained("facebook/blenderbot-400M-distill")
model = AutoModelForSeq2SeqLM.from_pretrained("facebook/blenderbot-400M-distill")

while True:
    print("Enter your prompt (or '1' to exit)\n")
    prompt = input()

    
    if prompt.isdigit() and int(prompt) == 1:
        print("Exiting the program...")
        break
    
    inputs = tokenizer.encode(prompt, return_tensors="pt")

    
    output = model.generate(inputs, max_length=100)

    response = tokenizer.decode(output[0], skip_special_tokens=True)
    
    print("Response:", response)
