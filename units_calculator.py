import tkinter as tk

def convert_km():
    km_per_hour = float(input_field.get().replace(',', '.'))
    meters_per_second = km_per_hour / 3.6
    mylable["text"] = meters_per_second
    mylable.update()
def convert_m():
    meters_per_second = float(input_field.get().replace(',', '.'))
    km_per_hour = meters_per_second * 3.6
    mylable["text"] =  km_per_hour
    mylable.update()
root = tk.Tk()
root.title ("Speed units calculator")
root.geometry("400x300")
root.resizable (width = False, height = False)
mylable = tk.Label(root, text ="Enter the speed")
mylable.place (x = 20, y = 60)
button_to_m = tk.Button(root, text ="Recalculate to m/s" , command = convert_m)
button_to_m.place (x = 30, y = 30)
button_to_km = tk.Button(root, text ="Recalculate to km/h" , command = convert_km)
button_to_km.place (x = 140, y = 30)
def is_digit(text):
    print(text)
    value = text.replace(',', '.')
    try:
        float(value)
        return True
    except vaue_error:
     return False


valid_value = (root.register(is_digit), "%P")
input_field = tk.Entry(root, validate = "key" , validatecommand = is_digit)
input_field.place (x = 30, y = 100)
root.mainloop()