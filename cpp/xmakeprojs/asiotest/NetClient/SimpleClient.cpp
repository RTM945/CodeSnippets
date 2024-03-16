#include "olc_net.h"

enum class CustomMsgType : uint32_t
{
    FireBullet,
    MovePlayer
};

int main()
{
    olc::net::message<CustomMsgType> msg;
    msg.header.id = CustomMsgType::FireBullet;

    int a = 1;
    bool b = true;
    float c = 3.14159f;

    struct
    {
        float x;
        float y;
    } d[5];

    msg << a << b << c << d;

    a = 99;
    b = false;
    c = 99.0f;
    
    msg >> d >> c >> b >> a;

    std::cout << "a = " << a << ", b = " << b << ", c = " << c << ",d = " << d << std::endl;

    return 0;
}