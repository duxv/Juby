y = ["Begin", "Break", "Catch", "Class", "Continue", "Data", "Define", "Do", "DynamicParam", "Else ", " Elseif ", " End ", " Enum ", " Exit ", " Filter ", " Finally ", " For ", " ForEach ", " From ", " Function ", " Hidden ", " If ", " In ", " InlineScript ", " Param ", " Process ", " Return ", " Static ", " Switch ", " Throw ", " Trap ", " Try ", " Until ", " Using ", " Var ", "While"]

z = []
for i in y:
    z.append(i.replace(" ", ""))

print(z)