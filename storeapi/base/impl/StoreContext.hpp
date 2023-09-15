#pragma once

/// Here lies the classes wrapping the MS API for testability.
/// Thus this code is inherently non-testable.

// For the underlying Store API
#include <winrt/windows.services.store.h>

// To provide the WinRT coroutine types.
#include <winrt/windows.foundation.h>

// To provide coroutines capable of returning more complex non-WinRT types.
#include <pplawait.h>

// For HWND and GUI-related Windows types.
#include <ShObjIdl.h>

#include <span>
#include <vector>

#include "../Purchase.hpp"

namespace StoreApi::impl {

// Wraps MS StoreContext type for testability purposes.
class StoreContext {
  winrt::Windows::Services::Store::StoreContext self =
      winrt::Windows::Services::Store::StoreContext::GetDefault();

 public:
  using Window = HWND;
  // Wraps MS StoreProduct type for testability purposes. This is not meant for
  // direct usage in high level code. The API is loose, the caller services must
  // tighten it up.
  class Product {
   public:
    Product(winrt::Windows::Services::Store::StoreProduct self) : self{self} {}
    // Whether the current user owns this product.
    bool IsInUserCollection() { return self.IsInUserCollection(); }

    // Assuming this is a Subscription add-on product the current user __owns__,
    // returns the expiration date of the current billing period.
    std::chrono::system_clock::time_point CurrentExpirationDate() const;

   protected:
    // Assuming this is a Subscription add-on product the current user __does
    // not own__, requests the runtime to display a purchase flow so users can
    // subscribe to this product. THis function returns early, the result will
    // eventually arrive through the supplied callback. This must be called from
    // a UI thread with the underlying store context initialized with the parent
    // GUI window because we need to render native dialogs. Thus, access to this
    // method must be protected so we can ensure it can only happen with GUI
    // clients, making API misuse harder to happen.
    // See
    // https://learn.microsoft.com/en-us/uwp/api/windows.services.store.storeproduct.requestpurchaseasync
    void PromptUserForPurchase(PurchaseCallback callback) const;

   private:
    winrt::Windows::Services::Store::StoreProduct self{nullptr};
  };

  // Returns a collection of products matching the supplied [kinds] and [ids].
  // Ids must match the Product IDs in Partner Center. Kinds can be:
  // Application; Game; Consumable; UnmanagedConsumable; Durable. See
  // https://learn.microsoft.com/en-us/uwp/api/windows.services.store.storeproduct.productkind#remarks
  std::vector<Product> GetProducts(std::span<const std::string> kinds,
                                   std::span<const std::string> ids);

  // Generates the user ID key (a.k.a the JWT) provided the server AAD [hToken]
  // and the [hUserId] the caller wants to have encoded in the JWT.
  winrt::Windows::Foundation::IAsyncOperation<winrt::hstring> GenerateUserJwt(
      winrt::hstring hToken, winrt::hstring hUserId) {
    return self.GetCustomerPurchaseIdAsync(hToken, hUserId);
  }

  // Initializes the GUI "subsystem" with the [parentWindow] handle so we can
  // render native dialogs, such as when purchase or other kinds of
  // authorization are required.
  void InitDialogs(Window parentWindow);
};

}  // namespace StoreApi::impl
