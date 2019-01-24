Rails.application.routes.draw do
  root to: 'home#index'

  post 'set_catalogs' => 'home#set_catalogs'

  get 'test_icon_url' => 'home#test_icon_url'

  get 'search' => 'home#search'

  devise_for :users, controllers: { registrations: 'users/registrations' }
  # For details onbu the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
