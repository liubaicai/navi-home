Rails.application.routes.draw do
  root to: 'home#index'

  post 'set_catalogs' => 'home#set_catalogs'

  devise_for :users, controllers: { registrations: 'users/registrations' }
  # For details onbu the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
