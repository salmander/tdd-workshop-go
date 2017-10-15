package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookServiceReturnsValidJSONWhenBookFound(t *testing.T) {
	drNoJSON := `{"isbn":"0099576929","ean":"1357924680","title":"Dr No","author":"Ian Fleming","price":3.99,"year_published":1993}`
	casinoRoyaleJSON := `{"isbn":"0230037496","ean":"2468013579","title":"Casino Royale","author":"Ian Fleming","price":4.99,"year_published":1990}`

	drNo := NewBook("0099576929", "1357924680", "Dr No", "Ian Fleming", 3.99, 1993)
	casinoRoyale := NewBook("0230037496", "2468013579", "Casino Royale", "Ian Fleming", 4.99, 1990)

	bookService := BookService{bookRepository}

	assert.Equal(t, drNoJSON, bookService.getBook("0099576929"))
	assert.Equal(t, casinoRoyaleJSON, bookService.getBook("0230037496"))
}

func TestBookServiceReutrnsValidJsonWhenBookNotFound(t *testing.T) {
	errorMoonrakerJSON := `{"status":404,"code":1,"title":"Book not found","detail":"Could not find book with ISBN: 0099576872"}`
	errorGoldfingerJSON := `{"status":404,"code":1,"title":"Book not found","detail":"Could not find book with ISBN: 1784871095"}`

	bookService := BookService{bookRepository}

	assert.Equal(t, errorMoonrakerJSON, bookService.getBook("0099576872"))
	assert.Equal(t, errorGoldfingerJSON, bookService.getBook("1784871095"))
}

func TestBookServiceReturnValidJsonWhenBookRepositoryUnavaiable(t *testing.T) {
	errorRepositoryJSON := `{"status":500,"code":2,"title":"Repository unavailable","detail":"Service unavailable, try again later"}`

	bookService := BookService{bookRepository}

	assert.Equal(t, errorRepositoryJSON, bookService.getBook("009957702X"))

}
