#pragma once

#include "net_common.h"
#include "net_tsqueue.h"
#include "net_message.h"


namespace olc 
{
    namespace net
    {
        template <typename T>
        class connection : public std::enable_shared_from_this<connection<T>>
        {
        public:
            enum class owner
            {
                server,
                client
            };

            connection(owner parent, asio::io_context& asioContext, asio::ip::tcp::socket socket, tsqueue<owned_message<T>>& qIn)
                        : m_asioContext(asioContext), m_socket(std::move(socket)), m_qMessagesIn(qIn)
            {
                m_nOwnerType = parent;
            }

            virtual ~connection()
            {}

            uint32_t GetID() const
            {
                return id;
            }

        public:
            void ConnectToClient(uint32_t uid = 0)
            {
                if (m_nOwnerType == owner::server)
                {
                    if (m_socket.is_open()) 
                    {
                        id = uid;
                    }
                }
            }
            bool ConnectToServer();
            bool Disconnect();
            bool isConnected() const
            {
                return m_socket.is_open();
            }

        public:
            bool Send(const message<T>& msg);

        protected:
            asio::ip::tcp::socket m_socket;
            asio::io_context& m_asioContext;

            tsqueue<message<T>> m_qMessagesOut;
            tsqueue<owned_message<T>>& m_qMessagesIn;
            owner m_nOwnerType = owner::server;
            uint32_t id = 0;
        };
    }
}