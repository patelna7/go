package main
import(
	"fmt"
	
)

func main(){
	const a float32 = 5.99 //Ціна 1 яблука
	const p int = 7	 //ціна 1 груші
	const s int = 23   //Загальна сума
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	fmt.Println("Віповідь:",9*a+8*float32(p),"грн.")
	fmt.Println("2. Скільки груш ми можемо купити?")
	fmt.Println("Віповідь: ",s/p)
	fmt.Println("3. Скільки яблук ми можемо купити?")
	fmt.Println("Віповідь: %d",float32(s)/a)
	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?")
	if ((float32(s)/(2*float32(p)+2*a))>1){
		fmt.Println("Віповідь: ТАК")
	}else{
		fmt.Println("Віповідь: НІ")
	}
}