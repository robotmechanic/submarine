# This file was generated by typhen-api

module TyphenApi::Model::Submarine
  class Room
    include Virtus.model(:strict => true)

    attribute :id, Integer, :required => true
    attribute :members, Array[User], :required => true
  end
end
