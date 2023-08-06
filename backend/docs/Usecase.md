# Commurz Usecases

This system will have 2 distinct modules:
- frontoffice
    - This module will be used by the customers.
- backoffice
    - This module will be used by the internal teams to manage the system.


## Front Office Usecases

### Register a User
- User is a guest that registered to the system.
- There are two types of user verified and unverified.
- Verified user is a user that has a valid email address.
- Unverified user is a user that has not confirmed his email address.

```mermaid
flowchart LR
    Start --> Guest[Guest]
    Guest -- register --> User::unverified
    User::unverified -- create confirm email --> Mailer
    Mailer -- send confirm email --> Guest
    Guest -- confirm email --> User::verified
    User::verified --> END
```

#### Add Product to Cart
- User can add multiple in stock products to his cart.
- If cart is not available it will create a new cart.
- User cannot add out of stock product to his cart.
- Adding or removing product from a cart does not affect the stock of the product.
- It will reserve the stock of the product when the user do a checkout.

```mermaid
flowchart LR
    Start --> User
    User -- add to cart --> Product
    Product --> CheckStok{is in stock?}
    CheckStok{is in stock?} -- yes: add to cart --> Cart
    Cart -- create if not exists --> Cart
    CheckStok{is in stock?} -- no: show error --> User
    Cart --> END
```

#### Remove Product from Cart
- User can remove a product from his cart.
- Removing a product from the cart does not affect the stock of the product.

```mermaid
flowchart LR
    Start --> User
    User -- remove product from cart --> Cart
    Cart -- remove --> Product
    Product --> END
```

#### Checkout
- Checkout is a creating an order from the cart.
- Checkout will reserve the stock of the product.
- Checkout will remove the cart.
- Order is a copy of the cart.
- The price of the order is the latest price of the product in the cart.
- The cart will be emptied after the checkout.
- A pending_payment invoice will be issued to the user after the checkout.
- A user can only checkout if no pending_payment invoice exists.

```mermaid
flowchart TD
    Start --> User
    User -- checkout --> Cart_check{is product in stock?}
    Cart_check{is product in stock?} -- yes --> Cart
    Cart_check{is product in stock?} -- no: show error --> User
    Cart -- checkout --> Order::pending_payment
    Order::pending_payment -- reserve stock --> Product
    Product -- empty cart --> Cart::emptied
    Cart::emptied -- create --> Invoice::pending_payment
    Invoice::pending_payment -- send to user --> Mailer
    Mailer --> END
```

#### Cancel Order
- User can cancel his order.
- Canceling an order will return the stock of the product.
- Canceling an order will invalidate the invoice.

```mermaid
flowchart TD
    Start --> User
    User -- cancel order --> Order
    Order -- return stock --> Product::restocked
    Product::restocked -- cancel order --> Order::canceled
    Order::canceled -- invalidate invoice --> Invoice::invalid
    Invoice::invalid --> END
```

#### Manual Pay an Invoice
- User can pay a pending invoice.
- Paying an invoice will mark it as wating_approval. 

```mermaid
flowchart TD
    Start --> User
    User -- pay invoice --> Invoice::pending_payment
    Invoice::pending_payment -- mark as wating_approval --> Invoice::wating_approval
    Invoice::wating_approval --> END
```

## Back Office Usecases

In the back office we will distinct the user that interact with the back office into several roles. Example for product related stuff will need a Product Admin role.

### Product

#### Manage Product
- Product Admin can add a product.

```mermaid
flowchart TD
    Start --> ProductAdmin
    ProductAdmin -- create/update/delete --> Product
    Product --> END
```

#### Add Product Stock
- Product Admin can increase the stock of a product.

```mermaid
flowchart TD
    Start --> ProductAdmin
    ProductAdmin -- add stock --> Product
    Product -- add stock --> ProductStock::increased
    ProductStock::increased --> END
```

### Reduce Product Stock
- Product Admin can reduce the stock of a product.

```mermaid
flowchart TD
    Start --> ProductAdmin
    ProductAdmin -- reduce stock --> Product
    Product -- reduce stock --> ProductStock::reduced
    ProductStock::reduced --> END
```

### Manual Payment

#### Approve Manual Payment
- Finance Admin can approve a manual payment.

```mermaid
flowchart TD
    Start --> FinanceAdmin
    FinanceAdmin -- approve --> Invoice::wating_approval
    Invoice::wating_approval -- mark as paid --> Invoice::paid
    Invoice::paid -- process order --> Order::pending_payment
    Order::pending_payment -- mark as processed --> Order::processed
    Order::processed --> END
```