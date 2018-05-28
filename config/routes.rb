Rails.application.routes.draw do
  root to: 'home#index'

  devise_for :users
  # For details onbu the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
