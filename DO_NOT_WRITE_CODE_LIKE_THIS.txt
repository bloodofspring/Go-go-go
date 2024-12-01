// ЭТОТ КОД РЕАЛЬНО РАБОТАЕТ, удачи :)
package main
import ("fmt"; "sort")
type Position struct {n int; checkTotal float64; price float64}
func (p Position) checkPrice() bool {return float64(p.n)*p.price == p.checkTotal}
func (p Position) repr() string {return fmt.Sprintf("Position{number: %d, price: %.2f, diff: %.2f}", p.n, p.price, p.checkTotal-float64(p.n)*p.price)}
type PositionSet struct {positions []Position}
func (ps PositionSet) total() float64 {var result float64; for _, p := range ps.positions {result += p.price * float64(p.n)}; return result}
// Len - Реализация сортировка в структуре PositionSet
func (ps PositionSet) Len() int {return len(ps.positions)}
// Swap - Реализация сортировка в структуре PositionSet
func (ps PositionSet) Swap(i, j int) {ps.positions[i], ps.positions[j] = ps.positions[j], ps.positions[i]}
// Less - Реализация сортировка в структуре PositionSet
func (ps PositionSet) Less(i, j int) bool {return ps.positions[i].price > ps.positions[j].price}
func ReadInput() (float64, []Position) {var numberOfProducts int; var totalPrice float64; fmt.Print("Enter total price: "); _, err := fmt.Scanf("%f", &totalPrice); fmt.Print("Enter number of products: "); _, err2 := fmt.Scanf("%d", &numberOfProducts); if err != nil || err2 != nil {return 0, []Position{}}; var price, n, total int; var products []Position; fmt.Println("Enter data:"); for i := 0; i < numberOfProducts; i++ {_, err = fmt.Scanf("%d *%d =%d", &price, &n, &total); if err != nil {panic(err)}; products = append(products, Position{n: n, checkTotal: float64(total), price: float64(price)})}; return totalPrice, products}
func checkPositions(checkPrice float64, products []Position) ([]Position, bool) {var incorrect []Position; for _, product := range products {if !product.checkPrice() {incorrect = append(incorrect, product)}}; sort.Sort(PositionSet{incorrect}); return incorrect, PositionSet{products}.total() == checkPrice}
func Start() {incorrectProducts, mainPriceCoincide := checkPositions(ReadInput()); if !mainPriceCoincide {fmt.Println("Итоговая цена и список продуктов не совпадают!")}; if len(incorrectProducts) == 0 {fmt.Println("No incorrect products!"); return}; fmt.Println("Incorrect products:"); for _, incorrectProduct := range incorrectProducts {fmt.Println(incorrectProduct.repr())}}
